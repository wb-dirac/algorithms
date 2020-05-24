package huffman

import (
	"fmt"
	"testing"
)

var str = []rune(`Sooner or later, it was bound to happen. On June 30, 1908, Moscow escaped destruction by three hours and four thousand kilometers—a margin invisibly small by the standards of the universe. On February 12, 1947, another Russian city had a still narrower escape, when the second great meteorite of the twentieth century detonated less than four hundred kilometers from Vladivostok, with an explosion rivaling that of the newly invented uranium bomb.
In those days there was nothing that men could do to protect themselves against the last random shots in the cosmic bombardment that had once scarred the face of the Moon. The meteorites of 1908 and 1947 had struck uninhabited wilderness; but by the end of the twenty-first century there was no region left on Earth that could be safely used for celestial target practice. The human race had spread from pole to pole. And so, inevitably….
At 0946 GMT on the morning of September 11 in the exceptionally beautiful summer of the year 2077, most of the inhabitants of Europe saw a dazzling fireball appear in the eastern sky. Within seconds it was brighter than the Sun, and as it moved across the heavens—at first in utter silence—it left behind it a churning column of dust and smoke.`)

func TestHuffmanTree_Encode(t *testing.T) {
	bin, tree := HEncode(str)
	Println(tree, []byte{})
	//fmt.Printf("%08b", bin)
	fmt.Printf("raw len: %d, encode len: %d, compression ratio: %.2f\n", len(str), len(bin), float64(len(bin))/float64(len(str))*100)
	raw := Decode(tree, bin)
	fmt.Printf("encode then decode: -----------\n%s\n----------------", string(raw))
	if string(str) != string(raw) {
		t.Errorf("encode then decode not match!!!")
	}
}
