

function setHeight(fieldId){
	var x = document.getElementsByClassName(fieldId);
	for (let v of x){
		v.style.height = v.scrollHeight+10+'px';
	}
}

function changeTheme(id){
	var mode = document.getElementById(id).value;
	if (mode == "Dark Mode") {
		document.getElementById(id).value = "Light Mode";
		document.documentElement.style.setProperty('--main-bg-color', '#fafafa');
		document.documentElement.style.setProperty('--nav-bg-color', '#343d46');
		document.documentElement.style.setProperty('--element-bg-color', 'white');
		document.documentElement.style.setProperty('--accent-color', '#00bfb3');
		document.documentElement.style.setProperty('--maint-text-color', 'white');
	} else {
		document.getElementById(id).value = "Dark Mode";
		document.documentElement.style.setProperty('--main-bg-color', '#3d3935');
		document.documentElement.style.setProperty('--nav-bg-color', 'white');
		document.documentElement.style.setProperty('--element-bg-color', '#fafafa');
		document.documentElement.style.setProperty('--accent-color', '#fbce40');
		document.documentElement.style.setProperty('--maint-text-color', 'black');
	}
	console.log(mode);
}

document.addEventListener("DOMContentLoaded", function() {
	setHeight('match');
});