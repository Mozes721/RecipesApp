import React, {useState, useEffect} from 'react';
import Searched from '../components/Searched'

const HomePage: React.FC = () => {
    const [text, setText] = useState("");
    const changeState = () => {
        console.log("FROM PARENT")
        console.log(text);
    }
  return (
      <section className="section py-4">
        <div className="container">
          <div className="column is-multiline is-centered">
            <div className="column has-text-centered">
              <span className="has-text-gray-dark">Welcome to My Recepie's</span>
              <h2 className="mt-2 mb-4 is-size-1 is-size-3-mobile has-text-weight-bold">Find recepies of your liking in outer pages and add them to your cooking wish list.</h2>
              <p className="subtitle has-text-gray mb-5">Otherwise search for anything particular you would like to cook bellow!</p>
              <Searched text={changeState} />
            </div>
          </div>
            <h1>{text}</h1>
        </div>
      </section>
  )
}

export default HomePage;