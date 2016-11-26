function Person(name){
  this.name = name;
  this.talk = function() {
    console.log("hi, " + this.name)
  }
}

const pat = new Person("pat");
pat.talk();
