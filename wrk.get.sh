#! /usr/bin/env bash

wrk -t50 -c100 -d20s http://0.0.0.0:5353/events/mytransactions?since=0a05e510-12ee-47b1-a7a7-2e4259d24809