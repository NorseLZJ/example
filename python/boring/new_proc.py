import subprocess
import ctypes

kernel32 = ctypes.WinDLL("kernel32.dll")
kernel32.CreateJobObjectW.argtypes = [ctypes.c_void_p, ctypes.c_wchar_p]
kernel32.CreateJobObjectW.restype = ctypes.c_void_p
kernel32.SetInformationJobObject.argtypes = [
    ctypes.c_void_p,
    ctypes.c_int,
    ctypes.c_void_p,
    ctypes.c_uint32,
]
kernel32.SetInformationJobObject.restype = ctypes.c_int
kernel32.AssignProcessToJobObject.argtypes = [ctypes.c_void_p, ctypes.c_void_p]
kernel32.AssignProcessToJobObject.restype = ctypes.c_int
kernel32.CloseHandle.argtypes = [ctypes.c_void_p]
kernel32.CloseHandle.restype = ctypes.c_int

JOBOBJECT_LIMIT_KILL_ON_JOB_CLOSE = 0x00002000
JOB_OBJECT_LIMIT_TIME = 0x00000004
JobObjectBasicLimitInformation = 2


class JOBOBJECT_BASIC_LIMIT_INFORMATION(ctypes.Structure):
    _fields_ = [("PerProcessUserTimeLimit", ctypes.c_int64)]


def create_job_object():
    return kernel32.CreateJobObjectW(None, None)


def set_job_object_limit_info(job_object, limit_flags, limit_info):
    info = JOBOBJECT_BASIC_LIMIT_INFORMATION()
    if JOB_OBJECT_LIMIT_TIME & limit_flags:
        info.PerProcessUserTimeLimit = limit_info
    kernel32.SetInformationJobObject(
        job_object,
        JobObjectBasicLimitInformation,
        ctypes.byref(info),
        ctypes.sizeof(info),
    )


def assign_process_to_job_object(job_object, process_id):
    kernel32.AssignProcessToJobObject(job_object, ctypes.c_void_p(process_id))


def close_handle(handle):
    kernel32.CloseHandle(handle)


def main():
    # 启动新的进程
    cmd = ["D:\\share\\fsv.exe", "-host", ":9999"]
    process = subprocess.Popen(cmd)

    # 创建一个新的 Job Object
    job_object = create_job_object()
    if job_object == 0:
        print("Error creating Job Object:", ctypes.WinError())
        return

    # 将新进程加入 Job Object
    assign_process_to_job_object(job_object, process.pid)

    # 设置 Job Object 的基本限制信息
    set_job_object_limit_info(
        job_object, JOB_OBJECT_LIMIT_TIME, int(0)
    )  # 设置用户模式执行时间限制，可根据需要调整

    print("New process started with PID:", process.pid)

    # 关闭 Job Object
    close_handle(job_object)


if __name__ == "__main__":
    main()
