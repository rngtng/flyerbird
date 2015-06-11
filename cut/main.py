#!/usr/bin/env python

import sys

from reportlab.pdfgen import canvas

c = canvas.Canvas(sys.stdout.buffer)
c.drawImage(sys.argv[1], 0,  200, 200, 200)
c.drawImage(sys.argv[2], 0, 400, 200, 200)
c.save()
