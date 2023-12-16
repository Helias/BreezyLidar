#include "gobreezylidar.h"
#include "URG04LX.hpp"

void breezy(void) {
  static const char *DEVICE = "/dev/ttyACM0";
  static const int NITER = 20;

  URG04LX laser;

  laser.connect(DEVICE);

  cout << "==============================================================="
       << endl;
  cout << laser << endl;
  cout << "==============================================================="
       << endl;

  for (int i = 1; i <= NITER; i++) {
    unsigned int data[1000];

    int ndata = laser.getScan(data);

    cout << "Iteration: " << i << ": ";

    if (ndata) {
      cout << "got " << ndata << " data points\n";
    } else {
      cout << "=== SCAN FAILED ===\n";
    }
  }
}

Breezylidar BreezylidarInit() { return NULL; }
