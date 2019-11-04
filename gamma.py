from dll import *

HDC = USER32.GetDC(None)

class gamma(Structure):
    _fields_ = [
        ("red", c_ushort * 256),
        ("green", c_ushort * 256),
        ("blue", c_ushort * 256)
    ]

def getGamma():
    g = gamma()
    GDI32.GetDeviceGammaRamp(HDC, byref(g))
    return g

def initGamma():
    g = getGamma()
    # 初始gamma设为0
    gamma = 0
    adjustGamma(g, gamma)
    return g, gamma

def adjustGamma(g, gamma):
    # 调节gamma
    for i in range(256):
        val = int(min(65535, max(0, (i / 256.0)**(4**gamma) * 65535 + 0.5)))
        g.red[i] = g.green[i] = g.blue[i] = val
    GDI32.SetDeviceGammaRamp(HDC, byref(g))
