import React, { Component } from 'react';
import {AppRegistry, Text } from 'react-native';
import { Container, Header, Title, Content, Button, Icon} from 'native-base';


export default class NativeBase extends Component {
  render() {
    return (
      <Container>

        <Header>
          <Button transparent>
            <Icon name='ios-arrow-back'  size={30} color="blue"/>
          </Button>
          <Title>Header</Title>
          <Button transparent>
            <Icon name="ios-menu" size={30} color="blue" />
          </Button>
        </Header>



        <Content>
          <Text> This is Content</Text>
        </Content>

      </Container>
    )
  }
}

AppRegistry.registerComponent('nativebase', () => NativeBase);
