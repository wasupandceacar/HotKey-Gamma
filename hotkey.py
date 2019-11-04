from gamma import *

import win32con
import ctypes
import ctypes.wintypes
import threading

GAMMAUP = False
GAMMADOWN = False

class Hotkey(threading.Thread):
    def run(self):
        global GAMMADOWN
        global GAMMAUP

        if not USER32.RegisterHotKey(None, 1, 0, 0xDD):
            print("Unable to register id1")

        if not USER32.RegisterHotKey(None, 2, 0, 0xDB):
            print("Unable to register id2")

        try:
            msg = ctypes.wintypes.MSG()
            while True:
                if USER32.GetMessageA(ctypes.byref(msg), None, 0, 0) != 0:

                    if msg.message == win32con.WM_HOTKEY:
                        if msg.wParam == 1:
                            GAMMAUP = True
                        elif msg.wParam == 2:
                            GAMMADOWN=True

                    USER32.TranslateMessage(ctypes.byref(msg))
                    USER32.DispatchMessageA(ctypes.byref(msg))

        finally:
            del msg
            USER32.UnregisterHotKey(None, 1)
            USER32.UnregisterHotKey(None, 2)

if __name__ == '__main__':
    hotkey = Hotkey()
    hotkey.start()

    g, gamma = initGamma()

    while True:
        if GAMMAUP == True:
            print("gamma up")
            gamma = max(-1, gamma - 0.05)
            adjustGamma(g, gamma)
            GAMMAUP = False
        elif GAMMADOWN == True:
            print("gamma down")
            gamma = min(1, gamma + 0.05)
            adjustGamma(g, gamma)
            GAMMADOWN = False

