# Title

After completion of TOC, remove the title line.

- [Title](#title)
  - [1. Check Box](#1-check-box)
  - [2. Code](#2-code)
  - [3. Image](#3-image)
  - [4. Link](#4-link)
  - [5. List](#5-list)
    - [5.1 Ordered List](#51-ordered-list)
    - [5.2 Unordered List](#52-unordered-list)
  - [6. Table](#6-table)
  - [7. Text](#7-text)

## 1. Check Box

<!-- Alt + C : x or not -->

- [x] hoge
- [ ] fuga

## 2. Code

- inline style `print 'hello'`
- block style

```ruby
print 'hello'
```

## 3. Image

When you export PDF, copy all the images in "images" folder  
and paste them in the same directory with the md file concerned.

![markdown](/images/markdown.png)

## 4. Link

<!-- select text + Cmd + V : quick generation -->

- [Markdown All in One](https://github.com/yzhang-gh/vscode-markdown#keyboard-shortcuts)
- [markdownlint](https://github.com/DavidAnson/markdownlint)
- [Markdown PDF](https://github.com/yzane/vscode-markdown-pdf/blob/master/README.ja.md)

## 5. List

### 5.1 Ordered List

1. hoge
2. fuga

### 5.2 Unordered List

- hoge
- fuga

## 6. Table

<!-- Alt + Sft + F : format table  -->

| L E F T | C E N T E R | R I G H T |
| :------ | :---------: | --------: |
| left    |   center    |     right |

## 7. Text

<!--
・Cmd + B : bold
・Cmd + I : italic
-->

- **bold**  
- *italic*  
- ***bold & italic***  
- ~~strikeout~~  
- <u>underline</u>  
- X<sup>2</sup>  
- X<sub>2</sub>

> blockquote
>> nest

## Omit Section <!-- omit in toc -->

- Cmd + P
  - \> Markdown All in one: Create Table of Contents
  - \> Markdown All in one: Update Table of Contents
  - \> Markdown PDF: Export (PDF)

- Cmd + Sft + V : Open Preview
- Cmd + K -> V  : Open Preview to the Side
- `<!-- omit in toc -->` : hide section title in TOC
- `<div style="page-break-before:always"></div>`  : insert page break to PDF
