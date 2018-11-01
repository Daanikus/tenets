package main

import "fmt"

type anyStruct struct {
    lock sync.Mutex
}

type bStruct struct {
    bLock sync.Mutex
}

type MockAddresses struct {
    lockA sync.Mutex
    bLock sync.Mutex
    lockC sync.Mutex
}

type AllGoodLocks struct {
    stateLock sync.Mutex
    stateBLock sync.Mutex
    anotherLock sync.Mutex
}

type BadLocks struct {
    lock sync.Mutex
    badlock sync.Mutex
    badlockC sync.Mutex
}
