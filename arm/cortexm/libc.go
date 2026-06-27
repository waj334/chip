package cortexm

import "unsafe"

//sigo:export __lock___libc_recursive_mutex
var lockLibcRecursiveMutex [1]byte

// init/close are no-ops (we use a global PRIMASK critical section, no per-lock state)
//
//sigo:export __retarget_lock_init
func retargetLockInit(lock unsafe.Pointer) {}

//sigo:export __retarget_lock_init_recursive
func retargetLockInitRecursive(lock unsafe.Pointer) {}

//sigo:export __retarget_lock_close
func retargetLockClose(lock unsafe.Pointer) {}

//sigo:export __retarget_lock_close_recursive
func retargetLockCloseRecursive(lock unsafe.Pointer) {}

// acquire/release: PRIMASK mask with nesting + saved state (as before)
var mallocLockNest uint32
var mallocLockPrimask uint32

//sigo:export __retarget_lock_acquire_recursive
func retargetLockAcquireRecursive(lock unsafe.Pointer) {
	state := DisableInterrupts()
	if mallocLockNest == 0 {
		mallocLockPrimask = state
	}
	mallocLockNest++
}

//sigo:export __retarget_lock_release_recursive
func retargetLockReleaseRecursive(lock unsafe.Pointer) {
	mallocLockNest--
	if mallocLockNest == 0 {
		EnableInterrupts(mallocLockPrimask)
	}
}

//sigo:export __retarget_lock_acquire
func retargetLockAcquire(lock unsafe.Pointer) { retargetLockAcquireRecursive(lock) }

//sigo:export __retarget_lock_release
func retargetLockRelease(lock unsafe.Pointer) { retargetLockReleaseRecursive(lock) }
