package fakeTML

import "math/rand"

// ShortUL short unordered list example
func ShortUL() string {
	return `<ul>
	<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
	<li>Aliquam tincidunt mauris eu risus.</li>
	<li>Vestibulum auctor dapibus neque.</li>
	</ul>`
}

// LongUL longer unordered list example
func LongUL() string {
	return `<ul>
<li>Morbi in sem quis dui placerat ornare. Pellentesque odio nisi, euismod in, pharetra a, ultricies in, diam. Sed arcu. Cras consequat.</li>
<li>Praesent dapibus, neque id cursus faucibus, tortor neque egestas augue, eu vulputate magna eros eu erat. Aliquam erat volutpat. Nam dui mi, tincidunt quis, accumsan porttitor, facilisis luctus, metus.</li>
<li>Phasellus ultrices nulla quis nibh. Quisque a lectus. Donec <strong>consectetuer</strong> ligula vulputate sem tristique cursus. Nam nulla quam, gravida non, commodo a, sodales sit amet, nisi.</li>
<li>Pellentesque fermentum dolor. Aliquam quam lectus, facilisis auctor, ultrices ut, elementum vulputate, nunc.</li>
</ul>
`
}

// ShortOL short ordered list example
func ShortOL() string {
	return `<ol>
<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
<li>Aliquam tincidunt mauris eu risus.</li>
<li>Vestibulum auctor dapibus neque.</li>
</ol>`
}

// Sentence html sentence wrapped in a p tag
func Sentence() string {
	return `<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.</p>`
}

// Paragraph html paragraph wrapped in a p tag
func Paragraph() string {
	return `<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum <i>tortor quam</i>, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>`
}

// LongParagraph longer html paragraph wrapped in a p tag
func LongParagraph() string {
	return `<p>Pellentesque habitant morbi <i>tristique</i> senectus et netus et malesuada fames ac turpis egestas. <b>Vestibulum</b> tortor quam, feugiat vitae, <strong>ultricies eget</strong>, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, commodo vitae, ornare sit amet, wisi. Aenean fermentum, elit eget tincidunt condimentum, eros ipsum rutrum orci, sagittis tempus lacus enim ac dui. Donec non enim in turpis pulvinar facilisis. Ut felis. Praesent dapibus, neque id cursus faucibus, tortor neque egestas augue, eu vulputate magna eros eu erat. Aliquam erat volutpat. Nam dui mi, tincidunt quis, accumsan porttitor, facilisis luctus, metus</p>`
}

// BlockQuote html blockquote
func BlockQuote() string {
	return `<blockquote><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus magna. Cras in mi at felis aliquet congue. Ut a est eget ligula molestie gravida. Curabitur massa. Donec eleifend, libero at sagittis mollis, tellus est malesuada tellus, at luctus turpis elit sit amet quam. Vivamus pretium ornare est.</p></blockquote>`
}

// Header h1 header
func Header() string {
	return `<h1>Header Level 1</h1>`
}

// HeaderTwo h2 header
func HeaderTwo() string {
	return `<h2>Header Level 2</h2>`
}

// HeaderThree h3 header
func HeaderThree() string {
	return `<h3>Header Level 3</h3>`
}

// HeaderFour h4 header
func HeaderFour() string {
	return `<h4>Header Level 4</h4>`
}

// HeaderFive h5 header
func HeaderFive() string {
	return `<h4>Header Level 4</h4>`
}

// Code html code example
func Code() string {
	return `<pre><code>
	#header h1 a {
		display: block;
		width: 300px;
		height: 80px;
	}
	</code></pre>`
}

// Form example - contains some classnames from bulma css
func Form() string {
	return `
	<form action="#" method="post">
    <div>
         <label for="name">Text Input:</label>
         <input type="text" name="name" id="name" value="" class="input" tabindex="1" />
    </div>

    <div>
         <h4>Radio Button Choice</h4>

         <label for="radio-choice-1">Choice 1</label>
         <input type="radio" name="radio-choice-1" id="radio-choice-1" tabindex="2" value="choice-1" />

     <label for="radio-choice-2">Choice 2</label>
         <input type="radio" name="radio-choice-2" id="radio-choice-2" tabindex="3" value="choice-2" />
    </div>

  <div>
    <label for="select-choice">Select Dropdown Choice:</label>
    <select name="select-choice" id="select-choice" class="select">
      <option value="Choice 1">Choice 1</option>
      <option value="Choice 2">Choice 2</option>
      <option value="Choice 3">Choice 3</option>
    </select>
  </div>

  <div>
    <label for="textarea" class="textarea">Textarea:</label>
    <textarea cols="40" rows="8" name="textarea" id="textarea"></textarea>
  </div>

  <div>
      <label for="checkbox">Checkbox:</label>
    <input type="checkbox" name="checkbox" id="checkbox" />
    </div>

  <div>
      <input type="submit" value="Submit" />
    </div>
</form>`
}

// Random returns a random chunk of html from all the options
func Random() string {
	str := ""
	r := rand.Intn(9)
	switch r {
	case 0:
		str += Paragraph()
	case 1:
		str += Sentence()
	case 2:
		ri := rand.Intn(1)
		if ri == 1 {
			str += Header()
		} else {
			str += HeaderTwo()
		}
	case 3:
		ri := rand.Intn(2)
		if ri == 1 {
			str += HeaderThree()
		} else if ri == 2 {
			str += HeaderFour()
		} else {
			str += HeaderFive()
		}
	case 4:
		str += BlockQuote()
	case 5:
		str += LongParagraph()
	case 6:
		str += ShortUL()
	case 7:
		str += ShortOL()
	case 8:
		str += LongUL()
	}
	return str
}

// Combo returns 5 combined chunks of random html seperated by a line break <br>
func Combo() string {
	str := ""
	upTo := rand.Intn(5)
	i := 0
	for i < upTo {
		str += Random() + "</br>"
		i++
	}
	return str
}

// ImageSquare returns a random sqaure image from lorempixel.com
func ImageSquare() string {
	return "//lorempixel.com/600/600"
}

// Image4by3 returns a random 4by3 image from lorempixel.com
func Image4by3() string {
	return "//lorempixel.com/1200/900"
}

// Image3by2 returns a random 3by2 image from lorempixel.com
func Image3by2() string {
	return "//lorempixel.com/1200/800"
}

// Image16by10 returns a random 16by10 image from lorempixel.com
func Image16by10() string {
	return "//lorempixel.com/1920/1200"
}

// Image16by9 returns a random 16by9 image from lorempixel.com
func Image16by9() string {
	return "//lorempixel.com/1920/1080"
}

// Image21by9 returns a random 21by9 image from lorempixel.com
func Image21by9() string {
	return "//lorempixel.com/1400/600"
}
