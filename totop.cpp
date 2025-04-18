#include <stdio.h>
#include <stdlib.h>
#include <X11/Xlib.h>

int main(int argc, char **argv)
{
if ( argc != 2 ) {
printf("Usage:\n\ttotop <window id>\n");
return 1;
}

Display *dsp = XOpenDisplay(NULL);

long id = strtol(argv[1], NULL, 16);
XRaiseWindow ( dsp, id );
XSetInputFocus ( dsp, id, RevertToNone, CurrentTime );

XCloseDisplay ( dsp );

return 0;
}
