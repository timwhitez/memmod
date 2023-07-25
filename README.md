Memmod
=======

Fork of Wireguard's memmod module

change virtualfree,virtualprotect,virtualalloc to Nt api Recycled Gate

## todo
参考项目:

https://github.com/Octoberfest7/Inline-Execute-PE/

https://github.com/timwhitez/Doge-MemX

1. 挂钩与命令行参数和退出进程相关的某些 API
2. runpe 传参
3. 捕获完整输出
4. 在内存加密上做一些思考
5. 敏感字符串替换
```
在当前进程下使用CreateThread加载一个解析过的PE文件，要完整捕获该PE的stdout和stderr输出，你可以尝试以下步骤：

使用CreatePipe函数创建两个匿名管道，一个用于捕获stdout输出，另一个用于捕获stderr输出。

将管道的写入端重定向到当前进程的标准输出和标准错误流。你可以使用SetStdHandle函数将GetStdHandle(STD_OUTPUT_HANDLE)和GetStdHandle(STD_ERROR_HANDLE)的返回值设置为管道的写入端。

在新线程中使用CreateProcess函数执行解析过的PE文件（这将成为新进程），并确保设置STARTUPINFO结构的hStdOutput和hStdError成员来允许子进程将输出写入到上一步中创建的管道。

在主线程中，使用ReadFile函数从管道的读取端读取子进程的stdout和stderr输出，直到读取结束。

下面是一个简单的示例代码，用于创建新线程并捕获stdout和stderr输出：
#include <windows.h>
#include <iostream>

HANDLE stdOutRead, stdOutWrite;
HANDLE stdErrRead, stdErrWrite;

DWORD WINAPI ThreadFunc(LPVOID lpParam) {
    STARTUPINFO si;
    PROCESS_INFORMATION pi;
    ZeroMemory(&si, sizeof(si));
    si.cb = sizeof(si);
    si.dwFlags = STARTF_USESTDHANDLES;
    si.hStdOutput = stdOutWrite;
    si.hStdError = stdErrWrite;

    if (CreateProcess("your_pe_file.exe", NULL, NULL, NULL, TRUE, 0, NULL, NULL, &si, &pi))
    {
        WaitForSingleObject(pi.hProcess, INFINITE);
        CloseHandle(pi.hProcess);
        CloseHandle(pi.hThread);
    }

    return 0;
}

int main() {
    SECURITY_ATTRIBUTES sa;
    sa.nLength = sizeof(SECURITY_ATTRIBUTES);
    sa.lpSecurityDescriptor = NULL;
    sa.bInheritHandle = TRUE;

    CreatePipe(&stdOutRead, &stdOutWrite, &sa, 0);
    CreatePipe(&stdErrRead, &stdErrWrite, &sa, 0);

    SetStdHandle(STD_OUTPUT_HANDLE, stdOutWrite);
    SetStdHandle(STD_ERROR_HANDLE, stdErrWrite);

    DWORD threadId;
    HANDLE hThread = CreateThread(NULL, 0, ThreadFunc, NULL, 0, &threadId);

    // 等待子线程结束
    WaitForSingleObject(hThread, INFINITE);

    // 从管道中读取子进程的stdout和stderr输出
    char buffer[4096];
    DWORD bytesRead;

    std::cout << "stdout:" << std::endl;
    while (ReadFile(stdOutRead, buffer, sizeof(buffer), &bytesRead, NULL) && bytesRead != 0) {
        std::cout.write(buffer, bytesRead);
    }

    std::cout << std::endl << "stderr:" << std::endl;
    while (ReadFile(stdErrRead, buffer, sizeof(buffer), &bytesRead, NULL) && bytesRead != 0) {
        std::cout.write(buffer, bytesRead);
    }

    CloseHandle(stdOutRead);
    CloseHandle(stdOutWrite);
    CloseHandle(stdErrRead);
    CloseHandle(stdErrWrite);

    return 0;
}

```


## Ref
[moloch--/memmod](https://github.com/moloch--/memmod)

