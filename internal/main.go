package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type ShortUID uint64

const (
	bitLenMachine  = 16
	bitLenTime     = 25
	bitLenSequence = 63 - bitLenMachine - bitLenMachine
	invalidUID     = ShortUID(1<<64 - 1)
	maxAge         = 1 << bitLenTime
	alphabet62     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	encodeURL      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

type Goflake struct {
	mtx                              sync.Mutex
	sequence, machineId, elapsedTime uint64
}

func New() *Goflake {
	machineId, _ := lower16BitPrivateIP()
	return &Goflake{mtx: sync.Mutex{}, machineId: uint64(machineId)}
}
func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
func (gf *Goflake) ID() ShortUID {
	gf.mtx.Lock()
	defer gf.mtx.Unlock()
	if now := uint64(time.Now().Unix()); gf.elapsedTime < now {
		gf.elapsedTime = now
		gf.sequence = 0

	} else {
		if gf.sequence >= (1 << bitLenSequence) {
			return invalidUID
		}
		gf.sequence++
	}

	return ShortUID((gf.sequence << (bitLenTime + bitLenMachine)) | ((gf.elapsedTime & uint64(1<<bitLenTime-1)) << bitLenMachine) | (gf.machineId))

}
func (uid ShortUID) ToBase62() string {

	s := make([]byte, 0, 9)
	l := ShortUID(len(alphabet62))
	for uid > 0 {
		s = append(s, alphabet62[uid%l])
		uid /= l
	}
	return string(s)

}
func (uid ShortUID) ToBase64() string {

	s := make([]byte, 0, 9)
	for uid > 0 {
		s = append(s, encodeURL[uid&63])
		uid = uid >> 6
	}
	return string(s)

}

func main() {
	gf := New()
	id := gf.ID()
	fmt.Println(gf.sequence, gf.elapsedTime&(1<<bitLenTime-1), id, id.ToBase64())
}
