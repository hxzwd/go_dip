#!/usr/bin/env python3.4

from matplotlib import pyplot as plt
import numpy as np
import sys

print("draw_hist.py\n\n")

if len(sys.argv) < 2:
	print("usage: ./draw_hist.py hist/baboon_hist.png\n")
	sys.exit(1)


file_name = sys.argv[1]
print("File name: ", file_name)

fd = open(file_name, "r")
hist = fd.readlines()
fd.close()

debug_tmp = None

if "co" in file_name and "[" in hist[0]:
	hist = map(lambda x: x.strip("[]\n").split(","), hist)
	debug_tmp = hist
#	hist = [ map(int, x) for x in hist ]
#	hist = np.array(hist)
	hist = hist[0:256]
	hist = np.array([ list(map(int, x)) for x in hist ])
	fig, ax = plt.subplots(nrows = 1, ncols = 1)
	color_plot = ax.pcolor(hist)
	ax.set_title("co-mat: {}".format(file_name))
	ax.grid(True)
	out_name = file_name.replace(".txt", ".png")
	print("save co_mat plot in {}".format(out_name))
	fig.savefig(out_name)
	plt.close(fig)
	sys.exit(0)

hist = map(lambda x: x.strip().split(" "), hist)
hist = [ tuple(map(int, x)) for x in hist ]

x, y = map(np.array, zip(*hist))
N = sum(y)

y_pdf = y/(N*1.0)

tmp = []
y = y.tolist()
for index, value in enumerate(y):
	tmp += [ index ]*value
num_bins = len(x)

out_name = file_name.replace(".txt", ".png")


plt.grid(True)
plt.title("hist: {}".format(file_name))
#plt.show()
fig, ax = plt.subplots(nrows = 2, ncols = 1)
ax[0].set_title("hist: {}".format(file_name))
ax[0].grid(True)
#ax[0].plot(x, y)
ax[0].hist(tmp, num_bins, facecolor = "green")
ax[1].set_title("pdf: {}".format(file_name))
ax[1].grid(True)
ax[1].plot(x, y_pdf)
print("save hist and pdf plot in {}".format(out_name))
fig.savefig(out_name)
plt.close(fig)





