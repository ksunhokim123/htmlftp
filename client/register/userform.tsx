import * as React from 'react';
import { Button, Container, Form, Header, Icon } from 'semantic-ui-react';

interface State {
  username: string;
  password: string;
}

interface Props {
  akey: string;
  next(endpoint: string): void;
}

export class UserForm extends React.Component<Props, State> {
  constructor(props: any) {
    super(props);
    this.state = {
      username: '',
      password: '',
    };
  }

  public onInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value,
    });
  }

  public onSumbit(e: any) {
    e.preventDefault();
    const data = {
      username: this.state.username,
      password: this.state.password,
      key: this.props.akey,
    };
    console.log(data)
    fetch('http://127.0.0.1:5353/api/users', {
          method: 'POST',
          headers: {
            'content-type': 'application/json',
          },
          body: JSON.stringify(data),
     })
     .then((response) => (response.status))
     .then((status) => {
       if (status === 200) {
         this.props.next(this.state.username);
       }
     });
  }

  public render() {
    return (
      <Form onSubmit={this.onSumbit.bind(this)}>
        <Form.Field>
          <label>Key</label>
          <input placeholder='username' name='username' onChange={this.onInputChange.bind(this)} />
          <input placeholder='passwrod' name='password' onChange={this.onInputChange.bind(this)} />
        </Form.Field>
        <Button fluid type='submit' onClick={this.onSumbit.bind(this)}>Submit</Button>
      </Form>
    );
  }
}
