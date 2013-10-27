pause-clementine-on-lock
========================

Pause [Clementine](http://www.clementine-player.org/) when the Windows session is locked.

Build
=====
    set CGO_LDFLAGS=-lwtsapi32
    
    # with powershell
    $env:CGO_LDFLAGS="-lwtsapi32";
    
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
