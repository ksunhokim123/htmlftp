import * as React from 'react';
import { Button, Container, Form, Header, Icon, Input } from 'semantic-ui-react';

interface State {
  key: string;
  error: boolean;
}

interface Props {
  next(key: string): void;
}

export class KeyInput extends React.Component<Props, State> {
  constructor(props: any) {
    super(props);
    this.state = {
      key: '',
      error: false,
    };
  }

  public onKeyChange(e: any) {
    this.setState(
      {
        key: e.target.value,
      }
    );
  }

  public onSumbit(e: any) {
    e.preventDefault();
    this.setState(
      {
        error: false,
      }
    );
    fetch('http://127.0.0.1:5353/api/keys/' + this.state.key, {
          method: 'GET',
          headers: {
            'content-type': 'application/json',
          },
     })
     .then((response) => (response.status))
     .then((status) => {
       if (status === 200) {
         this.props.next(this.state.key);
       } else if (status === 404) {
         this.setState(
           {
             error: true,
           }
         );
       }
     });

  }

  public render() {
    return (
      <Form onSubmit={this.onSumbit.bind(this)}>
        <Form.Group widths='equal'>
          <Form.Field name = 'key' control={Input}
          error={this.state.error ? true : false} label='key' placeholder='key'
          onChange={this.onKeyChange.bind(this)} />
        </Form.Group>
        <Button fluid type='submit' onClick={this.onSumbit.bind(this)}>Submit</Button>
      </Form>
    );
  }
}
