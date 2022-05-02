#!/bin/bash
echo "Runinng errcheck with the following command: 'errcheck -blank -asserts -ignore 'Write' cmd/*"
errcheck -blank -asserts -ignore 'Write' cmd/*
