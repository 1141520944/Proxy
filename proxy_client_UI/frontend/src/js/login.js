	// 动态流星雨特效代码
	window.onload = function () {
		var canvas = document.getElementById("canvas");
		canvas.width = window.innerWidth-20;
		canvas.height = window.innerHeight-20;
		var context = canvas.getContext("2d");

		// 定义流星的数量
		var count = 50;
		var meteors = [];

		// 定义流星对象
		function Meteor() {
		// 定义流星的位置和速度
			this.x = Math.random() * canvas.width;
			this.y = Math.random() * canvas.height;
			this.speed = Math.random() * 5 + 1;
			this.size = Math.random() * 2 + 1;

		// 绘制流星的函数
			this.draw = function () {
				context.beginPath();
				context.moveTo(this.x, this.y);
				context.lineTo(this.x - this.speed * 3, this.y + this.speed * 3);
				context.strokeStyle = "white";
				context.lineWidth = this.size;
				context.stroke();
			};

		// 更新流星的函数
		this.update = function () {
			this.x -= this.speed;
			this.y += this.speed;
			if (this.x < -50 || this.y > canvas.height +50) { // 修改这里
				this.x = Math.random() * canvas.width+400;
				this.y = -100; // 修改这里
				this.speed = Math.random() * 5 + 1;
				this.size = Math.random() * 2 + 1;
			}
			this.draw();
		};

        }

		// 初始化流星
		for (var i = 0; i < count; i++) {
			meteors.push(new Meteor());
		}

		// 绘制和更新流星
		function loop() {
			context.clearRect(0, 0, canvas.width, canvas.height);
			for (var i = 0; i < count; i++) {
				meteors[i].update();
			}
			requestAnimationFrame(loop);
		}
        loop();
	};
		function login() {
			var username = document.getElementById("username");
			var pass = document.getElementById("userpassword");
			if (username.value == "") {
				alert("请输入用户名");
			} else if(pass.value  == "") {
				alert("请输入密码");
				} else if(username.value == "admin" && pass.value == "666666"){
					window.location.href="navigation.html";
					} else {
						alert("请输入正确的用户名和密码！")
						}
		}
		function register(){
			window.location.href="register.html";
		}