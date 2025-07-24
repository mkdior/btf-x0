> Author: Hamza
> Description: Example implementation of a merkle tree service. Not intended for production use.

---

### Running the project

To get a feeling of how the packages/ services work, run and inspect the tests. The commit tree has two milestone tags; `PART.A` and `PART.B`. `PART.A` contains the main `merkle` package implementation, running `make` or `make test-pkgs` will automatically run the tests for you. To inspect the validity of the test cases, feel free to inspect the cases themselves and do hash calculations by yourself.

`PART.B` will contain both a docker image and a pre-built binary, so if you don't feel like installing go on your system, feel free to run the binary directly or build the docker image. The docker image however is designed to run in kubernetes.
