# Osgood File Server (OFS)

## Overview

The Osgood File Server (OFS) is designed to be a simple file server that utilizes protocol buffers (protobufs) and gRPC to transfer files. The server is written in Golang.

The server can be added to an existing gRPC server (by registering the filehandler via `filehandler.RegisterFileserviceServer`) or can be run as a stand-alone server using the `RunServer` function.
