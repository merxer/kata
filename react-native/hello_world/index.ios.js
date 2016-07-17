import React, { Component} from 'react';
import {
  AppRegistry,
  Text,
  StatusBar,
  View,
} from 'react-native';

class hello_world extends Component {
  render() {
    return (
      <View>
        <StatusBar hidden="true" />
        <Text> Hello World</Text>
      </View>
    );
  }
}

AppRegistry.registerComponent('hello_world', () => hello_world);
