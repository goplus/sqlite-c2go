package main

import unsafe "unsafe"

type struct_CCurHint struct {
}
type struct_Fts5Context struct {
}
type struct_Fts5Tokenizer struct {
}
type struct_KeyClass struct {
}
type struct_TreeView struct {
}
type struct___sFILEX struct {
}
type struct__filesec struct {
}
type struct_au_token struct {
}
type struct_knote struct {
}
type struct_label struct {
}
type struct_mount struct {
}
type struct_pgrp struct {
}
type struct_posix_cred struct {
}
type struct_proc struct {
}
type struct_proc_ident struct {
}
type struct_pthread_override_s struct {
}
type struct_session struct {
}
type struct_sigacts struct {
}
type struct_sqlite3_blob struct {
}
type struct_sqlite3_pcache struct {
}
type struct_sqlite3_stmt struct {
}
type struct_ucred struct {
}
type struct_user struct {
}
type struct_vnode struct {
}

func __builtin___memcpy_chk(unsafe.Pointer, unsafe.Pointer, uint, uint) unsafe.Pointer {
	panic("notimpl")
}
func __builtin___memmove_chk(unsafe.Pointer, unsafe.Pointer, uint, uint) unsafe.Pointer {
	panic("notimpl")
}
func __builtin___memset_chk(unsafe.Pointer, int32, uint, uint) unsafe.Pointer {
	panic("notimpl")
}
func __builtin___strlcat_chk(*int8, *int8, uint, uint) uint {
	panic("notimpl")
}
func __builtin___strlcpy_chk(*int8, *int8, uint, uint) uint {
	panic("notimpl")
}
func __builtin_bswap32(uint32) uint32 {
	panic("notimpl")
}
func __builtin_bswap64(uint64) uint64 {
	panic("notimpl")
}
func __builtin_fabs(float64) float64 {
	panic("notimpl")
}
func __builtin_fabsf(float32) float32 {
	panic("notimpl")
}
func __builtin_fabsl(float64) float64 {
	panic("notimpl")
}
func __builtin_huge_valf() float32 {
	panic("notimpl")
}
func __builtin_inf() float64 {
	panic("notimpl")
}
func __builtin_inff() float32 {
	panic("notimpl")
}
func __builtin_infl() float64 {
	panic("notimpl")
}
func __builtin_object_size(unsafe.Pointer, int32) uint {
	panic("notimpl")
}
func __darwin_check_fd_set_overflow(int32, unsafe.Pointer, int32) int32 {
	panic("notimpl")
}
func __error() *int32 {
	panic("notimpl")
}
func __sincos_stret(float64) struct___double2 {
	panic("notimpl")
}
func __sincosf_stret(float32) struct___float2 {
	panic("notimpl")
}
func __sincospi_stret(float64) struct___double2 {
	panic("notimpl")
}
func __sincospif_stret(float32) struct___float2 {
	panic("notimpl")
}
func __swbuf(int32, *struct___sFILE) int32 {
	panic("notimpl")
}
func __sync_synchronize() {
	panic("notimpl")
}
func access(*int8, int32) int32 {
	panic("notimpl")
}
func atoi(*int8) int32 {
	panic("notimpl")
}
func close(int32) int32 {
	panic("notimpl")
}
func confstr(int32, *int8, uint) uint {
	panic("notimpl")
}
func dlclose(__handle unsafe.Pointer) int32 {
	panic("notimpl")
}
func dlerror() *int8 {
	panic("notimpl")
}
func dlopen(__path *int8, __mode int32) unsafe.Pointer {
	panic("notimpl")
}
func dlsym(__handle unsafe.Pointer, __symbol *int8) unsafe.Pointer {
	panic("notimpl")
}
func fchmod(int32, uint16) int32 {
	panic("notimpl")
}
func fchown(int32, uint32, uint32) int32 {
	panic("notimpl")
}
func fcntl(int32, int32, ...interface {
}) int32 {
	panic("notimpl")
}
func flock(int32, int32) int32 {
	panic("notimpl")
}
func fprintf(*struct___sFILE, *int8, ...interface {
}) int32 {
	panic("notimpl")
}
func fsctl(*int8, uint, unsafe.Pointer, uint32) int32 {
	panic("notimpl")
}
func fstat(int32, *struct_stat) int32 {
	panic("notimpl")
}
func fstatfs(int32, *struct_statfs) int32 {
	panic("notimpl")
}
func fsync(int32) int32 {
	panic("notimpl")
}
func ftruncate(int32, int64) int32 {
	panic("notimpl")
}
func futimes(int32, *struct_timeval) int32 {
	panic("notimpl")
}
func getcwd(*int8, uint) *int8 {
	panic("notimpl")
}
func getenv(*int8) *int8 {
	panic("notimpl")
}
func geteuid() uint32 {
	panic("notimpl")
}
func gethostuuid(*uint8, *struct_timespec) int32 {
	panic("notimpl")
}
func getpid() int32 {
	panic("notimpl")
}
func gettimeofday(*struct_timeval, unsafe.Pointer) int32 {
	panic("notimpl")
}
func localtime(*int) *struct_tm {
	panic("notimpl")
}
func lstat(*int8, *struct_stat) int32 {
	panic("notimpl")
}
func malloc_create_zone(start_size uint, flags uint32) *struct__malloc_zone_t {
	panic("notimpl")
}
func malloc_default_zone() *struct__malloc_zone_t {
	panic("notimpl")
}
func malloc_set_zone_name(zone *struct__malloc_zone_t, name *int8) {
	panic("notimpl")
}
func malloc_size(ptr unsafe.Pointer) uint {
	panic("notimpl")
}
func malloc_zone_free(zone *struct__malloc_zone_t, ptr unsafe.Pointer) {
	panic("notimpl")
}
func malloc_zone_malloc(zone *struct__malloc_zone_t, size uint) unsafe.Pointer {
	panic("notimpl")
}
func malloc_zone_realloc(zone *struct__malloc_zone_t, ptr unsafe.Pointer, size uint) unsafe.Pointer {
	panic("notimpl")
}
func memcmp(__s1 unsafe.Pointer, __s2 unsafe.Pointer, __n uint) int32 {
	panic("notimpl")
}
func mkdir(*int8, uint16) int32 {
	panic("notimpl")
}
func mmap(unsafe.Pointer, uint, int32, int32, int32, int64) unsafe.Pointer {
	panic("notimpl")
}
func munmap(unsafe.Pointer, uint) int32 {
	panic("notimpl")
}
func open(*int8, int32, ...interface {
}) int32 {
	panic("notimpl")
}
func pread(__fd int32, __buf unsafe.Pointer, __nbyte uint, __offset int64) int {
	panic("notimpl")
}
func pthread_create(**struct__opaque_pthread_t, *struct__opaque_pthread_attr_t, func(unsafe.Pointer) unsafe.Pointer, unsafe.Pointer) int32 {
	panic("notimpl")
}
func pthread_join(*struct__opaque_pthread_t, *unsafe.Pointer) int32 {
	panic("notimpl")
}
func pthread_mutex_destroy(*struct__opaque_pthread_mutex_t) int32 {
	panic("notimpl")
}
func pthread_mutex_init(*struct__opaque_pthread_mutex_t, *struct__opaque_pthread_mutexattr_t) int32 {
	panic("notimpl")
}
func pthread_mutex_lock(*struct__opaque_pthread_mutex_t) int32 {
	panic("notimpl")
}
func pthread_mutex_trylock(*struct__opaque_pthread_mutex_t) int32 {
	panic("notimpl")
}
func pthread_mutex_unlock(*struct__opaque_pthread_mutex_t) int32 {
	panic("notimpl")
}
func pthread_mutexattr_destroy(*struct__opaque_pthread_mutexattr_t) int32 {
	panic("notimpl")
}
func pthread_mutexattr_init(*struct__opaque_pthread_mutexattr_t) int32 {
	panic("notimpl")
}
func pthread_mutexattr_settype(*struct__opaque_pthread_mutexattr_t, int32) int32 {
	panic("notimpl")
}
func pwrite(__fd int32, __buf unsafe.Pointer, __nbyte uint, __offset int64) int {
	panic("notimpl")
}
func random() int {
	panic("notimpl")
}
func read(int32, unsafe.Pointer, uint) int {
	panic("notimpl")
}
func readlink(*int8, *int8, uint) int {
	panic("notimpl")
}
func rename(__old *int8, __new *int8) int32 {
	panic("notimpl")
}
func rmdir(*int8) int32 {
	panic("notimpl")
}
func sleep(uint32) uint32 {
	panic("notimpl")
}
func srandomdev() {
	panic("notimpl")
}
func stat(*int8, *struct_stat) int32 {
	panic("notimpl")
}
func statfs(*int8, *struct_statfs) int32 {
	panic("notimpl")
}
func strcmp(__s1 *int8, __s2 *int8) int32 {
	panic("notimpl")
}
func strcspn(__s *int8, __charset *int8) uint {
	panic("notimpl")
}
func strlen(__s *int8) uint {
	panic("notimpl")
}
func strncmp(__s1 *int8, __s2 *int8, __n uint) int32 {
	panic("notimpl")
}
func strrchr(__s *int8, __c int32) *int8 {
	panic("notimpl")
}
func sysconf(int32) int {
	panic("notimpl")
}
func sysctlbyname(*int8, unsafe.Pointer, *uint, unsafe.Pointer, uint) int32 {
	panic("notimpl")
}
func time(*int) int {
	panic("notimpl")
}
func unlink(*int8) int32 {
	panic("notimpl")
}
func utimes(*int8, *struct_timeval) int32 {
	panic("notimpl")
}
func write(__fd int32, __buf unsafe.Pointer, __nbyte uint) int {
	panic("notimpl")
}
