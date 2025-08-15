onkeydown = (e) =>
	e.key == "e" && !e.shiftKey && !e.metaKey && !e.ctrlKey && !e.altKey
		? (checkExpressed.checked ^= true, e.preventDefault())
		: 0;
