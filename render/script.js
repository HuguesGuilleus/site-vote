onkeydown = (e) =>
	e.key == "e" &&
		e.target == document.body &&
		!e.shiftKey &&
		!e.metaKey &&
		!e.ctrlKey &&
		!e.altKey
		? (checkExpressed.checked ^= !e.preventDefault())
		: 0;

S = Array.from(
	document.querySelectorAll("ul.sub>li"),
	(i) => [i, i.innerText.toLowerCase()],
);
(search || 0).oninput = () => {
	V = search.value.toLowerCase();
	S.map(([i, t]) => i.hidden = !t.includes(V));
};
