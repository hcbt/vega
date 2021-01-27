import os
import sys
import warnings
warnings.filterwarnings("ignore")

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import scipy.io.wavfile

# Load audio into numpy array, 
# waveform y, sampling rate sr
def load_file(filename):
    sr, y = scipy.io.wavfile.read(filename)

    return y, sr

# Audio to spectrogram
def build_spectrogram(filename, id):
    y, sr = load_file(filename)

    plt.figure(figsize=(15, 10))
    ax = plt.axes()
    ax.set_axis_off()
    plt.specgram(y[:, 0], NFFT = 2048, Fs=2, Fc=0, noverlap=128, cmap=plt.get_cmap("inferno"), sides="default", mode="default", scale="dB")
    plt.savefig(id + ".png", bbox_inches='tight', transparent=True, pad_inches=0.0)
    plt.close()

def main():
    build_spectrogram(sys.argv[1], sys.argv[2])

if __name__ == "__main__":
    main()