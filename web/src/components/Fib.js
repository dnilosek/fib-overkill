import React, { Component } from 'react';
import Jubmotron from 'react-bootstrap/Jumbotron'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import axios from 'axios';

class Fib extends Component {
  state = {
    seenIndexes: [],
    values: {},
    index: ''
  };

  componentDidMount() {
    this.fetchValues();
    this.fetchIndexes();
  }

  async fetchValues() {
    const values = await axios.get('/api/values/current');
    this.setState({ values: values.data });
  }

  async fetchIndexes() {
    const seenIndexes = await axios.get('/api/values/all');
    this.setState({
      seenIndexes: seenIndexes.data
    });
  }

  handleSubmit = async event => {
    event.preventDefault();

    await axios.post('/api/values', {
      index: this.state.index
    });
    this.setState({ index: '' });
    this.fetchValues();
    this.fetchIndexes();
  };

  renderSeenIndexes() {
    return this.state.seenIndexes.map(number => number).join(', ');
  }

  renderValues() {
    const entries = [];
    // eslint-disable-next-line
    for (let key in this.state.values) {
      entries.push(
        <div key={key}>
          For index {key} I calculated {this.state.values[key]}
        </div>
      );
    }

    return entries;
  }

  render() {
    return (
      <div>
        <Jubmotron>
          <Form onSubmit={this.handleSubmit} className="form-inline justify-content-center">
              <Form.Label className="my-1 mr-2">Enter your index: </Form.Label>
              <Form.Control
                className="my-1 mr-sm-2"
                value={this.state.index}
                onChange={event => this.setState({ index: event.target.value })}
              />
              <Button className="my-1" type="submit">Submit</Button>
          </Form>
        </Jubmotron>
        <Container>
          <Row>
            <Col>
              <h1>Indexes I have seen:</h1>
              {this.renderSeenIndexes()}
            </Col>
            <Col>
              <h1>Calculated Values:</h1>
              {this.renderValues()}
            </Col>
          </Row>
        </Container>
      </div>
    );
  }
}

export default Fib;
