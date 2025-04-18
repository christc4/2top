# 2top

Provided is my Go rewrite

To compile run c++ totop.cpp -L/usr/X11R6/lib -lX11 -o totop

The shell script, id=$(xwininfo -name "st" | grep id: | awk "{ print \$4 }") && totop $id

Alternatively you can ommit the second line, and only use the third with the id if you know it

Credit to original cpp [author](https://www.linuxquestions.org/questions/linux-general-1/how-to-bring-up-application-window-to-front-from-shell-script-83545/)
