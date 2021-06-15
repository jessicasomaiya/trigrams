# Trigram Predictive Text

A Golang algorithm that uses trigrams to create predictive text, given training data.

### Run

To run the program, run the command:

`make run`

The default port is :8080.

### Learn

To train the algorithm using a POST request on `/learn`:

`curl --header "Content-Type: text/plain" --data-binary @test_text.txt localhost:8080/learn`

For example, to train the algorithm with The Wonderful Wizard of Oz:

`curl --header "Content-Type: text/plain" --data-binary @training-data/wiz.txt localhost:8080/learn`

More examples can be found in `training-data`. Credit to [Project Gutenberg](https://www.gutenberg.org/))

### Generate

To generate predictive text, use a GET request on `/generate`:

`curl http://localhost:8080/generate`
