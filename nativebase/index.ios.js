import React, { Component } from 'react';
import {AppRegistry, Text} from 'react-native';
import { Container, Header, Title, Content } from 'native-base';

export default class NativeBase extends Component {
  render() {
    return (
      <Container>
        <Header>
          <Title>Header Na</Title>
        </Header>

        <Content>
          <Text> This is content section </Text>
        </Content>
      </Container>
    )
  }
}

AppRegistry.registerComponent('nativebase', () => NativeBase);
