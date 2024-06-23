package server

import (
	"crypto/tls"
	"time"

	"github.com/thomas-osgood/OGOR/output"
	"github.com/thomas-osgood/ofs/ofsauthenticators"
	"github.com/thomas-osgood/ofs/protobufs/filehandler"
	"github.com/thomas-osgood/ofs/server/internal/ofstypes"
	"google.golang.org/grpc"
)

// fileserver object. this can be added to an existing
// grpc server by registering it with the filehandler
// service via "filehandler.RegisterFileserviceServer".
type FServer struct {
	filehandler.UnimplementedFileserviceServer

	authenticator ofsauthenticators.OFSAuthenticator
	debug         bool
	downloadsdir  string
	encryptor     OFSEncryptor
	printer       *output.Outputter
	rootdir       string
	transferCfg   ofstypes.TransferConfig
	uploadsdir    string
}

// object used to set the confiruation options for
// a new fileserver.
type FServerOption struct {
	// authenticator to use while processing requests.
	Authenticator ofsauthenticators.OFSAuthenticator
	// if this is set to true, output will be printed to
	// STDOUT while the server is running.
	Debug bool
	// subdirectory within the rootdir where uploaded
	// files will be saved.
	Downloadsdir string
	// encryptor the fileserver will use.
	Encryptor OFSEncryptor
	// number of maximum concurrent downloads the server is
	// allowed to have active at one time.
	MaxDownloads int
	// number of maximum concurrent uploads the server is
	// allowed to have active at one time.
	MaxUploads int
	// fileserver root directory.
	Rootdir string
	// maximum allowed time a file upload or download can take.
	TransferTimeout time.Duration
	// subdirectory within the rootdir where files that
	// can be downloaded to a client are stored.
	Uploadsdir string
}

type GrpcConfig struct {
	listenaddr string
	listenport int
	options    []grpc.ServerOption
}

type GrpcConfigOpt struct {
	Listenaddr string
	Listenport int
	Options    []grpc.ServerOption
	TLSCert    tls.Certificate
}
