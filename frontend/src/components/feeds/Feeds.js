import React, { useEffect, useState } from "react";
import "./Feeds.css";

var url = "http://localhost:8080/feeds";


function Feeds() {
    const [loading, setLoading] = useState(true);
    const [feeds, setFeeds] = useState([]);
    const renderHTML = (rawHTML) => React.createElement("div", { dangerouslySetInnerHTML: { __html: rawHTML } });


    const fetchFeeds = async () => {
        setLoading(true)

        try {
            const response = await fetch(url);
            const jsonFeeds = await response.json();
            setFeeds(jsonFeeds);
            setLoading(false);
            console.log(feeds)
        } catch (error) {
            setLoading(true);
            console.log(error);
        }
    }

    useEffect(() => {
        fetchFeeds();
    }, []);

    if (loading) {
        return <main>
            <center><h1>Is Loading</h1></center>
        </main>
    }

    return (
        <main>
            {
                feeds.map(entry => {
                     return <div key={entry.Channel.id}>
                        <h1>{entry.Channel.Title}</h1>
                        <p>{entry.Channel.Description}</p>
                        {
                            entry.Channel.Item.map(item => {
                                return <div className="feed">
                                    <a href={item.Link}>{item.Title}</a>
                                    <p>{renderHTML(item.Description)}</p>
                                </div>
                            })
                        }
                    </div>
                })
            }

        </main>
    )
}


export default Feeds;