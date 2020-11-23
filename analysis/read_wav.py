import scipy.io.wavfile
import sys
import numpy
numpy.set_printoptions(threshold=sys.maxsize)

def main():
    sr, y = scipy.io.wavfile.read("../tests/kiQtEDrdYDY.wav")

    print(y)

if __name__ == "__main__":
    main()