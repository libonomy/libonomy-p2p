
Thank You for visiting our open source project.

Libonomy is one of a kind blockchain which through innovation and creativity has achieved
all the set goals with great success. The kind of innovation that attracts and reaches out to
masses of people who have knowledge and understanding that this system is fulfilling its role
of greatness.
Many only talk about the next level blockchain technology that would fix all the existing
problems all in one, Libonomy is already on a road to fulfill this ideology.

We are designing and coding an autonomous blockchain that will run on first ever AI based consensus algorithm, which will make
libonomy worlds first every autonomous, interoperable and scalable blockchain.
To learn more about Libonomy visit [https://libonomy.com](https://libonomy.com).

To learn more about the libonomy upcoming blockchain [Read this](https://libonomy.com/assets/pdf/yellow-paper.pdf).

### Benefits
The major difference is that Libonomy Blockchain does not use any of the previous consensus
algorithms that have been used for a very long time. All of these consensus algorithms have
their own drawbacks, Libonomy believes in providing an error-free consensus engine that has
been architecture very carefully. It uses Artificial Intelligence – automated, computer
generated engine that saves time, energy and gives high throughput, is scalable, interoperable and autonomous.

### Project Status
We are working hard towards our first patch - which is public testnet running libonomy AI consensus algorithm.
Current version is the simulation of the p2p protocol that is being developed  by the libonomy developers. In the upcoming
days our team will be release the full node running the consensus algithm with which users will be able to carry out the transactions as well.

### Important
This repository is currently under development and from time to time the file structure is meant to be updated. On releasing
further features and repositories we will be announcing on our social media platforms and on website. 

### How to begin

```bash
git clone git@github.com:libonomy/libonomy-p2p.git
```
_-- or --_

Fork the project from https://github.com/libonomy/libonomy-p2p

Go 1.14's Modules are being used in the package so  it is best to place the code **outside** your `$GOPATH`. 

### Setting Up Local Environment

Install [Go 1.14 or later](https://golang.org/dl/) for your platform, if you haven't already.

Ensure that `$GOPATH` is set correctly and that the `$GOPATH/bin` directory appears in `$PATH`.

Before building we need to install `protoc` (ProtoBuf compiler) and some tools required to generate ProtoBufs. 
```bash
make install
```
This will invoke `setup_env.sh` which supports Linux . 
The configurations for other platforms will be released soon.



### Building Simulator
To make a build of p2p simulator, from the project root directory, use:
```
make p2p-build
```

This will (re-)generate protobuf files and build the `libonomoy-p2p` binary, saving it in the `build/` directory.


Platform-specific binaries are saved to the `/build` directory.

### Connection Establishment
In order to connect to the p2p-simulater you need to provide the parameters for the network configuration which is explained below.
The configuration file is named as sim_config.toml which you can find in the master branch.

#### Simulater Configuration

1. Navigate to build directory under libonomy-p2p after generating build of p2p simulater.
2. Start the simulater with the command given below

```bash
./p2p-simulate --tcp-port 7152 --config ./sim_config.toml -d ./libo_data
```
As soon as you are done you will be connected to the p2p network.

#### Upcoming Release
- In the coming month we will be releasing the p2p light node for the community.
- AI data extraction algorithm (minimal)
- Wallet implementation in JS
- Testnet Faucets
- .....
