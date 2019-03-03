#!/bin/bash

docker run --rm -v $(pwd):/documents/ asciidoctor/docker-asciidoctor asciidoctor-pdf -r asciidoctor-diagram README.adoc
