class Person {
  constructor (name) {
    this.name = name
  }
  talk() {
    console.log('Hi, ' + this.name)
  }
}

const pat = new Person('pat');
pat.talk();
