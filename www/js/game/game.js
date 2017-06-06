'use strict'

class Game {
	constructor(contanerName) {
		log('constructor')

	   this.pg = new Phaser.Game(
	      800, 600,
	      Phaser.AUTO,
	      contanerName,
			{
				preload: ()=> this.preload(),
				create: ()=> this.create(),
				update: ()=> this.update(),
				render: ()=> this.render()
			}
	   );

		this.player = null;
	}

	initNetwork() {
		this.socket = io.connect(window.location.host, {path: "/ws/", transports: ['websocket']});

		this.socket.on('tick', this.onTick);
		this.socket.on('playerConnected', this.onPlayerConnected);
		this.socket.on('playerDisconnected', this.onPlayerDisconnected);
	}

	preload() {
	   this.pg.load.image('sky', 'assets/sky.png');
	   this.pg.load.image('ground', 'assets/platform.png');
	   this.pg.load.image('star', 'assets/star.png');
	   this.pg.load.spritesheet('dude', 'assets/dude.png', 32, 48);
	}

	create() {
	   this.pg.add.sprite(0, 0, 'sky'); //  A simple background for our game
		this.player = new Player(this.pg);
		this.pg.input.onDown.add(this.onInputDown, this);

		this.initNetwork()

		log("game created")
	}

	update() {
		this.player.update();
	}

	render() {
	   this.pg.debug.cameraInfo(this.pg.camera, 8, 16);
	}

	onInputDown(pointer) {
		// pointer will contain the pointer that activated this event
		let msg = {
			x: pointer.x,
			y: pointer.y
		}
		log('click to ' + JSON.stringify(msg))
		this.socket.emit("move", JSON.stringify(msg))
	}

	onTick(msg) {
		//log("tick " + msg);
	}

	onPlayerConnected(msg) {
		log("connected player, id: " + msg);
	}

	onPlayerDisconnected(msg) {
		log("disconnected player, id: " + msg);
	}
}
