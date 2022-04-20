var DoOnceActive = true;
var login_loader;

function createWinsockAlert(message) {
	var tmpDiv = $('<div/>')
    .addClass("alert alert-info")
    tmpDiv.append(message);
	$("#main_box_alert").append(tmpDiv);
	scrollToBottom();
}

function createNormalAlert(message) {
	var tmpDiv = $('<div/>')
    .addClass("alert alert-secondary")
    tmpDiv.append(message);
	$("#main_box_alert").append(tmpDiv);
	scrollToBottom();
}

function createGoodAlert(message) {
	var tmpDiv = $('<div/>')
    .addClass("alert alert-success")
    tmpDiv.append(message);
	$("#main_box_alert").append(tmpDiv);
	scrollToBottom();
}

function createErrorAlert(message) {
    var tmpDiv = $('<div/>')
    .addClass("alert alert-danger")
    tmpDiv.append(message);
	$("#main_box_alert").append(tmpDiv);
	scrollToBottom();
}

function scrollToBottom() {
	const messages = document.getElementById('main_box_alert');
	if (messages != null)
		messages.scrollTop = messages.scrollHeight;
}

function switchLogin() {
	createNormalAlert("Enter Erupe ID and Password, then press [Log in]");
	login_loader = document.getElementById("login_load");
	var usernameSaved = localStorage.getItem('username');
	var passwordSaved = localStorage.getItem('password');
	var checkboxSaved = localStorage.getItem('saveLogin');

	if (usernameSaved != "null") {
		document.getElementById("username").value = usernameSaved;
	}
	if (passwordSaved != "null") {
		document.getElementById("password").value = passwordSaved;
	}
	if (checkboxSaved != "null") {
		if (checkboxSaved == "true") {
			var checkbox = document.getElementById('saveLogin');
			checkbox.checked = true;
		} else {
			localStorage.removeItem('username');
			localStorage.removeItem('password');
			localStorage.removeItem('saveLogin');
		}
	}
	$("#saveLoginText").on("click", function(e) {
		if ($("#saveLogin").prop("checked")) {
			$("#saveLogin").prop("checked", false);
		} else {
			$("#saveLogin").prop("checked", true);
		}
	});
  $("#login_form").submit(function(e) {
		e.preventDefault();
    username = $("#username").val();
    password = $("#password").val();
		if (username == "") {
			createErrorAlert("Please enter Erupe ID!");
		} else if (password == "") {
			createErrorAlert("Please enter Password!");
		} else {
			login_loader.style.display = "block";
			createNormalAlert("Authenticating...");
			try {
				window.external.loginCog(username, password, password);
			} catch(e) {
				createErrorAlert("Error on loginCog: " + e + ".");
			}
			checkAuthResult();
		}
  });
}

$(function() {
  $("#main_titlebar").on("click", function(e) {
      window.external.beginDrag(true);
  });
  $("#main_exit").on("click", function(e) {
    window.external.closeWindow();
  });
	$("#main_reduce").on("click", function(e) {
    window.external.minimizeWindow();
	});
  $(window).on("message onmessage", function(e) {
    var data = e.originalEvent.data;
		CheckMessage(data);
  });
	$("#main_config_button").click(function(){
    try {
      window.external.openMhlConfig();
    } catch(e) {
      createErrorAlert("Error on openMhlConfig: " + e + ".");
    }
  });
  doLauncherInitalize();
	switchLogin();
});

function switchCharsel() {
	$('#main_parent').empty();
	$('#main_parent').append(charselHtml);
	try {
		var charInfo = window.external.getCharacterInfo();
	} catch (e) {
		createErrorAlert("Error getting character info!");
	}
	try {
		$xmlDoc = new ActiveXObject("Microsoft.XMLDOM");
		$xmlDoc.async = "false";
		$xmlDoc.loadXML(charInfo);
		$xml = $($xmlDoc);
	} catch (e) {
		createErrorAlert("Error parsing character info XML!" + e);
	}

	try {
		$($xml).find("Character").each(function () {
		  createCharListItem(
			$(this).attr('name'),
			$(this).attr('uid'),
			$(this).attr('weapon'),
			$(this).attr('HR'),
			$(this).attr('GR'),
			$(this).attr('lastLogin'),
			$(this).attr('sex')
			);
		});
	} catch (e) {
		createErrorAlert("Error searching character info xml!" + e);
	}

	$(".charsel_char").click(function () {
		if (!$(this).hasClass("active")) {
		  $(".charsel_char.active").removeClass("active");
		  $(this).addClass("active");
		}
	});

  $(function() {
    var selectedUid = $(".charsel_char.active").attr("uid");
  	$("#charsel_new_char").on("click", function(e) {
  		alert("Not yet implemented");
  	});
  	$("#charsel_logout").on("click", function(e) {
			$('#main_parent').empty();
			createErrorAlert("Disconnected.");
			$('#main_parent').append(loginHtml);
			switchLogin();
  	});
  });

  $("#charsel_launch").on("click", function () {
  	try {
  		elementID = parent.document.getElementById("main_start_now");
  		elementID.style.display = "block";
  	} catch(e) {
  		alert(e);
  	}
    var selectedUid = $(".charsel_char.active").attr("uid");
    try {
      window.external.selectCharacter(selectedUid, selectedUid)
    } catch (e) {
    	createErrorAlert("Error on selectCharacter: " + e + ".");
      try {
    		elementID = parent.document.getElementById("main_start_now");
    		elementID.style.display = "none";
    	} catch(e) {
    		alert(e);
    	}
    }
    setTimeout(function () {
      window.external.exitLauncher();
    }, 3000);
  });
}

function saveAccount() {
	var username = document.getElementById('username').value;
	var password = document.getElementById('password').value;
	var checkbox = document.getElementById('saveLogin');

	if (checkbox.checked == true){
		localStorage.setItem('username', username);
		localStorage.setItem('password', password);
		localStorage.setItem('saveLogin', 'true');
	} else {
		localStorage.removeItem('username');
		localStorage.removeItem('password');
		localStorage.removeItem('saveLogin');
	}
}

function checkAuthResult() {
  var loginResult = window.external.getLastAuthResult();
  console.log('|' + loginResult + '|');
  if (loginResult == "AUTH_PROGRESS") {
    setTimeout(checkAuthResult, 10);
  } else if (loginResult == "AUTH_SUCCESS") {
		saveAccount();
		createGoodAlert("Connected.");
		createNormalAlert("After selecting a character, press [Launch]");
		switchCharsel();
  } else {
		login_loader.style.display = "none";
    createErrorAlert("Error logging in!");
  }
}

function playKey(e) {
  var audio = new Audio("./audio/sys_cursor.mp3");
  audio.play();
}

function createCharListItem(name, uid, weapon, HR, GR, lastLogin, sex) {
	var icon;
	const unixTimestamp = lastLogin;
	const milliseconds = unixTimestamp * 1000;
	const dateObject = new Date(milliseconds);
	lastLogin = dateObject.toLocaleDateString("en-US");
	lastLoginString = "";
	for (var i = 0; i < lastLogin.length; i++) {
		if (lastLogin[i] != "‎") { // invisible LTR char
			lastLoginString += lastLogin[i];
		}
	}
	if (sex == "M"){
		sex = "♂";
	} else {
		sex = "♀";
	}
	if (HR > 999) {
		HR = 999;
	}
	if (GR > 999) {
		GR = 999;
	}

	if (weapon == "片手剣") {
		weapon = "Sword & Shield";
		icon = "./resources/icons/SS.png";
	} else if (weapon == "双剣") {
		weapon = "Dual Swords";
		icon = "./resources/icons/DS.png";
	} else if (weapon == "大剣") {
		weapon = "Great Sword";
		icon = "./resources/icons/GS.png";
	} else if (weapon == "太刀") {
		weapon = "Long Sword";
		icon = "./resources/icons/LS.png";
	} else if (weapon == "ハンマー") {
		weapon = "Hammer";
		icon = "./resources/icons/H.png";
	} else if (weapon == "狩猟笛") {
		weapon = "Hunting Horn";
		icon = "./resources/icons/HH.png";
	} else if (weapon == "ランス") {
		weapon = "Lance";
		icon = "./resources/icons/L.png";
	} else if (weapon == "ガンランス") {
		weapon = "Gunlance";
		icon = "./resources/icons/GL.png";
	} else if (weapon == "穿龍棍") {
		weapon = "Tonfa";
		icon = "./resources/icons/T.png";
	} else if (weapon == "スラッシュアックスF") {
		weapon = "Switch Axe F";
		icon = "./resources/icons/SAF.png";
	} else if (weapon == "マグネットスパイク") {
		weapon = "Magnet Spike";
		icon = "./resources/icons/MS.png";
	} else if (weapon == "ヘビィボウガン") {
		weapon = "Heavy Bowgun";
		icon = "./resources/icons/HS.png";
	} else if (weapon == "ライトボウガン") {
		weapon = "Light Bowgun";
		icon = "./resources/icons/LB.png";
	} else if (weapon == "弓") {
		weapon = "Bow";
		icon = "./resources/icons/B.png";
	} else {
	weapon = "Unknown"
		icon = "./resources/icons/null.png";
	}

	if (DoOnceActive) {
		DoOnceActive = false;
		var topDiv = $('<div/>')
		.attr("href", "#")
		.attr("uid", uid)
		.addClass("charsel_char list-group-item list-group-item-action flex-column align-items-start active");
	} else {
		var topDiv = $('<div/>')
		.attr("href", "#")
		.attr("uid", uid)
		.addClass("charsel_char list-group-item list-group-item-action flex-column align-items-start");
	}
	var topLine = $('<div/>')
	.addClass("char_name")
	.append($('<h1/>').addClass("mb-1").text(name)
	);
	var bottomLine = $('<div/>')
	.addClass("char_info")
	.append($('<div id="icon_weapon"/>').prepend($('<img>',{id:'theImg',src:icon})))
	.append($('<div id="weapon_title"/>').text('Current Weapon'))
	.append($('<div id="weapon_name"/>').text(weapon))
	.append($('<div id="hr_lvl"/>').text('HR' + HR))
	.append($('<div id="gr_lvl"/>').text('GR' + GR))
	.append($('<div id="sex"/>').text(sex))
	.append($('<div id="uid"/>').text('ID: ' + uid))
	.append($('<div id="lastlogin"/>').text('Last Login: ' + lastLoginString));
	topDiv.append(topLine);
	topDiv.append(bottomLine);
	$("#charsel_list").append(topDiv);
}

function CheckMessage(message){
	// Good Alert
	if (message == "Connected."){
		createGoodAlert(message);
	}
	// Normal Alert
	else if (message == "Authenticating..."){
		createNormalAlert(message);
	}
	else if (message == "Select or create a character, then press [Launch]"){
		createNormalAlert(message);
	}
	// Error Alert
	else {
		createErrorAlert(message);
	}
}

function doLauncherInitalize() {
	createWinsockAlert("Winsock Ver. [2.2]");
  try {
    window.external.getMhfMutexNumber();
  } catch(e) {
    createErrorAlert("Error on getMhfMutexNumber: " + e + ".");
  }
  try {
    var serverListXml = window.external.getServerListXml();
  } catch(e) {
    createErrorAlert("Error on getServerListXml: " + e + ".");
  }
  if (serverListXml == "") {
    createErrorAlert("Server list is empty!");
  }
  console.log(serverListXml);
  try {
    var lastServerIndex = window.external.getIniLastServerIndex();
  } catch(e) {
    createErrorAlert("Error on getIniLastServerIndex: " + e + ".");
  }
  console.log("Last server index:" + lastServerIndex);
  try {
    window.external.setIniLastServerIndex(0);
  } catch(e) {
    createErrorAlert("Error on setIniLastServerIndex: " + e + ".");
  }
  try {
    var mhfBootMode = window.external.getMhfBootMode();
  } catch(e) {
    createErrorAlert("Error on getMhfBootMode: " + e + ".");
  }
  console.log("mhfBootMode:" + mhfBootMode);
  try {
    var userId = window.external.getUserId();
  } catch(e) {
    createErrorAlert("Error on getUserId: " + e + ".");
  }
  console.log("userId:" + userId);
  try {
    var password = window.external.getPassword();
  } catch(e) {
    createErrorAlert("Error on getPassword: " + e + ".");
  }
  console.log("password:" + password);
	loginHtml = "<div class='container'>\
		<div class='row'>\
			<div class='col-1' style='max-width: 3.5%;'></div>\
			<div class='col-6'>\
				<div class='segment'>\
					<form id='login_form'>\
						<h5 class='mb-1' style='font-size: 16px;color: White; position: relative; left: 5px;' ><b>Erupe ID</b></h5>\
						<div class='form-group'>\
							<input type='text' class='form-control' id='username' placeholder='Username' onkeyup='playKey(event);' onclick='CUE_Selected()' onmouseover='CUE_Cursor()' value='' autocomplete='on'>\
						</div>\
						<h5 class='mb-1' style='font-size: 16px;color: White; position: relative; left: 5px;' ><b>Password</b></h5>\
						<div class='form-group'>\
							<input type='password' class='form-control' id='password' placeholder='Password' onkeyup='playKey(event);' onclick='CUE_Selected()' onmouseover='CUE_Cursor()' value='' autocomplete='on'>\
						</div>\
						<h5 class='mb-1' style='font-size: 16px;color: White; position: relative; left: 5px;' ><b>Server Selection</b></h5>\
						<div class='form-group'>\
							<input type='text' class='form-control' id='ServerName' value='Erupe' readonly>\
						</div>\
						<button id='login_submit' type='submit' class='btn btn-primary' onclick='CUE_Confirm();' onmouseover='CUE_Cursor()'>Log In</button>\
					</form>\
					<div id='login_save_login' >\
						<input type='checkbox' id='saveLogin' name='saveLogin' style='position: relative; left: 5px; cursor: pointer;'>\
						<label for='scales' id='saveLoginText' style='position: relative; left: 5px; bottom: 1.58px; cursor: pointer;'>Save Login Details</label>\
					</div>\
				</div>\
			</div>\
		</div>\
	</div>\
	<div id='login_load' style='display: none;'>\
		<img src='resources/load.gif' style='width: 32px;position: absolute;left: 190px;top: 125px;'>\
	</div>"
	charselHtml = "<div style='height: 30px;'></div>\
	<div class='container' style='height: 250px;'>\
		<div class='row'>\
			<div class='col-12' >\
				<div id='charsel_list' class='list-group'></div>\
			</div>\
		</div>\
		<div class='row' style='position: relative; top: 30px;'>\
			<div class='col-12' style='height: 100%;'>\
				<button id='charsel_logout' class='btn btn-primary' onclick='CUE_Selected();' onmouseover='CUE_Cursor()'>Log Out</button>\
				<button id='charsel_new_char' class='btn btn-primary' onclick='CUE_Selected()' onmouseover='CUE_Cursor()'>Add New Character</button>\
				<button id='charsel_launch' class='btn btn-primary' onclick='CUE_Starting();' onmouseover='CUE_Cursor()'>Launch</button>\
			</div>\
		</div>\
		<button id='charsel_char_list_up' class='btn btn-primary' onclick='CUE_Selected();' onmouseover='CUE_Cursor()'></button>\
		<button id='charsel_char_list_down' class='btn btn-primary' onclick='CUE_Selected();' onmouseover='CUE_Cursor()'></button>\
	</div>"
	$('#main_parent').append(loginHtml);
}





