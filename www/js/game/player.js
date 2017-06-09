'use strict'

class Player {
	constructor(phaserGame) {
		this.val = 10;
		this.targetX = 100;
		this.sprite = phaserGame.add.sprite(100, 450, 'car');
		this.sprite.anchor.set(0.5);
	}

	setTargetDirection(x) {
		this.targetX = x;
	}

	update() {
		if (this.sprite.x != this.targetX) {
			let d = this.targetX - this.sprite.x;
			this.sprite.x += Math.sign(d) * Math.min(Math.abs(d), this.val);
		}
	}
}
