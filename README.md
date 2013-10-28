pause-clementine-on-lock
========================

Pause [Clementine](http://www.clementine-player.org/) when the Windows session is locked using the [remote control feature](https://code.google.com/p/clementine-player/wiki/RemoteControl).

Download
========
- [pause-clementine-on-lock-32.7z](https://s3.amazonaws.com/pause-clementine-on-lock/pause-clementine-on-lock-32.7z) 32-bit (594 KB)
- [pause-clementine-on-lock-64.7z](https://s3.amazonaws.com/pause-clementine-on-lock/pause-clementine-on-lock-64.7z) 64-bit (606.1 KB)

Build
=====
To build you need [go1.1.2](http://golang.org/) and [MinGW](http://www.mingw.org/) (I use [MingW-64](http://mingw-w64.sourceforge.net/download.php#mingw-builds)).

Ensure that `gcc.exe` is in your `PATH` environment variable.

    # windowsgui prevents a console window from showing (remove to debug)
    go build -ldflags -H=windowsgui

Usage
=====
Rename config file to `config.gcfg` and edit as needed.

Possibles values for `onLock` and `onUnLock` are:
- PLAY
- PAUSE
- STOP
- NOTHING

The process run hidden so if you want to stop it, use the process manager and look for `pause-clementine-on-lock.exe`.
