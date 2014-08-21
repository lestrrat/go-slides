static void
child_handler(int sig) {
    pid_t pid;
    int status;

    // UNIXなプロセスモデルであればプロセス終了をwait()関数で
    // 待てるし、その終了ステータスも取得できる
    // …が、これはgoroutineではできない
    while ((pid = waitpid(-1, &status, WNOHANG)) > 0) {
       ...
    }
}


