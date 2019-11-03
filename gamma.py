from dll import *

HDC = USER32.GetDC(None)

class gamma(Structure):
    _fields_ = [
        ("red", c_ushort * 256),
        ("green", c_ushort * 256),
        ("blue", c_ushort * 256)
    ]

def initGamma():
    g = gamma()
    GDI32.GetDeviceGammaRamp(HDC, byref(g))
    # 初始亮度设为128
    b = 128
    adjustGamma(g, b)
    return g, b

def adjustGamma(g, brightness):
    # 调节gamma
    for i in range(256):
        val = min(65535, (128 + brightness) * i)
        g.red[i] = g.green[i] = g.blue[i] = val
    GDI32.SetDeviceGammaRamp(HDC, byref(g))
