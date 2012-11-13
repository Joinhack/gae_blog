function vector(x,y) {
	this.x = x;
	this.y = y;
}
vector.prototype = {
	add: function(v) {
		return new vector(this.x + v.x,this.y + v.y);
	},
	sub: function(v) {
		return new vector(this.x - v.x,this.y - v.y);
	},
	mul: function(v) {
		return new vector(this.x * v,this.y * v);
	},
	div: function(v) {
		return new vector(this.x / v.x,this.y / v.y);
	},
	neg: function(v) {
		return new vector(-this.x,-this.y);
	}
}


function start(func) {
	var timer = null;
	loop = function () {
		if(timer)
			window.clearTimeout(timer);
		if(func())
			timer = window.setTimeout(loop,30);
	}
	loop();
}

function Color(r,g,b) {
	this.r = r;
	this.g = g;
	this.b = b;
}


function Particle(pos,vel,accel,size,life,color) {
	this.position(pos);
	this.velocity(vel);
	this.life(life);
	this.age(0);
	this.size(size);
	this.color(color);
	this.accel(accel);
}

Particle.prototype = {
	_attr: function(n,v) {
		if(v != null) {
			this["p_" + n] = v;
		} else {
			return this['p_' + n];
		}
	},
	velocity: function(v) {
		return this._attr("velocity",v);
	},
	position: function(p) {
		return this._attr('position',p);
	},
	color: function(c) {
		return this._attr('color',c);
	},
	accel: function(a) {
		return this._attr('accel',a);
	},
	life: function(l) {
		return this._attr('life',l);
	},
	age: function(a) {
		return this._attr('age',a);
	},
	size: function(s) {
		return this._attr('size',s);
	},
	move: function(t) {
		pos = this.position();
		accel = this.accel();
		vel = this.velocity();
		this.velocity(vel.add(accel.mul(t)));
		this.position(pos.add(vel.mul(t)));
		this.age(this.age() + t);
	},
	draw: function(ctx) {
		color = this.color();
		r = Math.floor(color.r * 255);
		g = Math.floor(color.g * 255);
		b = Math.floor(color.b * 255);
		pos = this.position();
		a = (1 -this.age()/this.life()).toFixed(2);
		ctx.fillStyle = "rgba("+r+","+g+","+b+","+a+")";
		ctx.beginPath();
		ctx.arc(pos.x,pos.y,this.size(),0,Math.PI*2,true);
		ctx.closePath();
		ctx.fill();
	}
};

function Sense1() {
}


t = 0.1;

Sense1.prototype = {
	children: [],
	addChild: function(c) {
		this.children.push(c);
	},
	run: function(ctx,w,h) {
		nc = [];
		ctx.fillStyle='rgba(0,0,0,0.22)';
		ctx.fillRect(0,0,w,h);

		for(i = 0; i < this.children.length; i++) {
			var c = this.children[i];
			pos = c.position();
			if(pos.x < 0 || pos.x > w) { 
				v = c.velocity();
				v.x = - v.x;
				c.velocity(v);
			}
			if(pos.y < 0 || pos.y > h) {
				v = c.velocity();
				v.y = - v.y
				c.velocity(v);
			}
			c.move(t);
			c.draw(ctx);
			if(c.age() < c.life())
				nc.push(c);
			else
				delete c;
		}
		delete this.children;
		this.children = nc;
	}
};


function randDirect() {
    var a = Math.random() * 2 * Math.PI;
    return new vector(Math.cos(a), Math.sin(a));
}

function randDirect2() {
    var a = Math.random() * 2 * Math.PI;
    return new vector(Math.cos(a), a);
}

function randColor() {
	i = Math.random()*3;
	r = i%0||255;
	g = i%1||255;
	b = i%2||255;
	return new Color(r,g,b);
}

window.onload = function(){
	cv1 = document.getElementById("cv");
	ctx = cv1.getContext("2d");
	sens = new Sense1();


	w  = cv1.width;
	h = cv1.height;

	start((function(ctx){
		return function(){
			o1 = new Particle(new vector(200,150),
				randDirect().mul(15),new vector(0,9),
				5,8,randColor());

			sens.addChild(o1);
			sens.run(ctx,w,h);
			return true;
		};
	})(ctx));

}
