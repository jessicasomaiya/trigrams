# Trigram Learning

A Golang algorithm that uses trigram driven methods to create predictive text, given training text.

### General

- There is a brain which is trained using the /learn endpoint
- An output is generated and outputted using the /generate endpoint

### Run

To run the server run:

`make run`

The default port is :8080.

### Learn

To train the algorithm using a POST request, use:

`curl --data-binary @test_text.txt localhost:8080/learn`

For example, to train the algorithm using Wizard of Oz:

`curl --data-binary @training-data/wiz.txt localhost:8080/learn`

More can be found in `training-data`. Credit to [Project Gutenberg](https://www.gutenberg.org/))

### Generate

To generate predictive text, use a GET request on `/generate`:

`curl http://localhost:8080/generate`
