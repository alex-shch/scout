'use strict'

class ClickPoint {
   constructor(game, x, y) {
      this.game = game;
      this.sprite = game.pg.add.sprite(x, y, 'star');
      this.sprite.anchor.set(0.5)
   }

   update() {
      this.sprite.y += 5;
      if (this.sprite.y > 610) {
         this.game.delPoint(this);
         this.sprite.destroy();
      }
   }
}
