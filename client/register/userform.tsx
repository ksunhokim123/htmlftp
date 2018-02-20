import * as React from 'react';
import { Button, Container, Form, Header, Icon, Input } from 'semantic-ui-react';

interface State {
  username: string;
  password: string;
  error1: boolean;
  error2: boolean;
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
      error1: false,
      error2: false,
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
    const regex = /^[a-zA-Z0-9]+$/;
    let ok = true;
    let state = {
      error1: false,
      error2: false,
    };
    if (!regex.test(this.state.username)) {
      state.error1 = true;
      ok = false;
    }
    if (!regex.test(this.state.password)) {
      state.error2 = true;
      ok = false;
    }
    this.setState(state);
    if (ok) {
      const data = {
        username: this.state.username,
        password: this.state.password,
        key: this.props.akey,
      };
      fetch('http://127.0.0.1:5353/api/users', {
            method: 'POST',
            headers: {
              'content-type': 'application/json',
            },
            body: JSON.stringify(data),
       })
       .then((response) => (response.status))
       .then((status) => {
         if (status === 201) {
           this.props.next(this.state.username);
         } else {
           this.setState(
             {
               error1: true,
               error2: true,
             }
           );
         }
       });
    }

  }

  public render() {
    return (
      <Form onSubmit={this.onSumbit.bind(this)}>
        <Form.Input fluid name = 'username' control={Input}
        error={this.state.error1 ? true : false} label='username' placeholder='username'
        onChange={this.onInputChange.bind(this)} />
        <Form.Input fluid name = 'password' control={Input}
        error={this.state.error2 ? true : false} label='password' placeholder='password'
        onChange={this.onInputChange.bind(this)} type='password'/>
        <Button fluid type='submit' onClick={this.onSumbit.bind(this)}>Submit</Button>
      </Form>
    );
  }
}
