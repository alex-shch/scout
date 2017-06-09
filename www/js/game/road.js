'use strict'

class Road {
   constructor(phaserGame) {
      let g = phaserGame.add.group();

      let u = phaserGame.add.sprite(0, 0, 'road');
      u.scale.x = 380.0 / u.width;
      u.scale.y = 600.0 / u.height;
      u.y = -u.height;
      //this.sprite = s;
      let v = phaserGame.add.sprite(0, 0, 'road');
      v.scale.x = 380.0 / v.width;
      v.scale.y = 600.0 / v.height;
      v.y = u.height;


      this.sprites = [u, v];

   }

   update() {
      let u = this.sprites[0];
      let v = this.sprites[1];

      if (u.y >= 0) {
         v.y = u.y - u.height;
         this.sprites[0] = v;
         this.sprites[1] = u;
      }

      u.y += 5;
      v.y += 5;
   }
}
