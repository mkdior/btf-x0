> Author: Hamza
> Description: Example implementation of a merkle tree service. Not intended for production use.

---

### Running the project

To get a feeling of how the packages/ services work, run and inspect the tests. The commit tree has two milestone tags; `PART.A` and `PART.B`. `PART.A` contains the main `merkle` package implementation, running `make` or `make test` will automatically run the tests for you. To inspect the validity of the test cases, feel free to inspect the cases themselves and do hash calculations by yourself.

`PART.B` will contain a server implementation on which you can test out adding users/ generating merkle trees and proofs. To run the "happy-flow" of PART.B, first run `make run`, which builds all binaries, and runs the `server` binary on `127.0.0.1:8082`. Following that, run `make run-requests`. This will automatically run all requests for you. Of course, if this somehow breaks, please inspect the makefile, and run the CURL commands manually.
