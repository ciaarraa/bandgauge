# bandgauge

## Why does this exist?
To measure internet speed and record the result. Intended to be run on a schedule.

## Prerequisites
1. You must have (iperf3)[https://iperf.fr/iperf-download.php] downloaded and available on your path
2. You must have a remote iperf3 server running

## Configure
Run `bandwidth configure` to configure the IP and port of the remote iperf3 server.

## Usage
`bandwidth test` to run a test and output the results to the default file