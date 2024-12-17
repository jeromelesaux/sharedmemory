package model

import (
	"fmt"

	"github.com/gen2brain/shm"
)

type SharedMemory interface {
	Set(buf []byte) error
	Get() ([]byte, error)
}

const SHMSIZE = 1024 * 1000

type InMemory struct {
	ptrID int
}

// NewInMemory create a shared memory structure
// and alloc the memory in memory
func NewInMemory(ptrID int) (InMemory, error) {
	i := InMemory{
		ptrID: ptrID,
	}

	return i, i.Init()
}

// Init allocate the memory zone to be shared
func (m *InMemory) Init() error {
	if m.ptrID != 0 {
		return nil
	}
	var err error
	m.ptrID, err = shm.Get(shm.IPC_PRIVATE, SHMSIZE, shm.IPC_CREAT|0777)
	return err
}

// Close flush the memory allocation
func (m *InMemory) Close() error {
	return shm.Rm(m.ptrID)
}

// Set save in memory the slice of bytes
func (m *InMemory) Set(buf []byte) error {
	data, err := shm.At(m.ptrID, 0, 0)
	if err != nil {
		return err
	}
	copy(data[:], buf[:])

	return shm.Dt(data)
}

// Get retreive the memory and returns the slice of bytes
func (m *InMemory) Get() ([]byte, error) {
	return shm.At(m.ptrID, 0, 0)
}

func (m InMemory) String() string {
	return fmt.Sprintf("PtrID :%d", m.ptrID)
}

func (m InMemory) ID() int {
	return m.ptrID
}
