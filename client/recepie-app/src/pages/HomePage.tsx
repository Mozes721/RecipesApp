import React, {useState, useEffect} from 'react';
import Searched from '../components/searched/Searched';
import SearchedResult from '../components/searched/SearchedResult';
const HomePage: React.FC = () => {
    const [text, setText] = useState<string | number>("");
    const callBack = (val: string | number) => {
        setText(val);
        console.log(val);
        alert(val);
    }
  return (
      <section className="section py-4">
        <div className="container">
          <div className="column is-multiline is-centered">
            <div className="column has-text-centered">
              <span className="has-text-gray-dark">Welcome to My Recepie's</span>
              <h2 className="mt-2 mb-4 is-size-1 is-size-3-mobile has-text-weight-bold">Find recepies of your liking.</h2>
              <p className="subtitle has-text-gray mb-5">Otherwise search for  particular you would like to cook bellow!</p>
              <Searched text={callBack} />
            </div>
          </div>
            <SearchedResult  text={text} />
        </div>
      </section>
  )
}

export default HomePage;