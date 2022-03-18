(ns racegex.websocket
  (:require [aleph.http :as http]
            [compojure.core :as compojure :refer [GET]]
            [compojure.route :as route]
            [manifold.bus :as bus]
            [manifold.deferred :as d]
            [manifold.stream :as s]
            [ring.middleware.params :as params]
	    [clojure.data.json :as json]))

(def non-websocket-request
  {:status 400
   :headers {"content-type" "application/text"}
   :body "Expected a websocket request."})

(def users-per-room (atom {}))

(defn insert-user
  [users room user]
    (let
      [room-users (find users room)]
      (if room-users
        (assoc users room (assoc (val room-users) user "online"))
	(assoc users room {user "online"}))))

(def rooms (bus/event-bus))

(defn join-handler
  [socket room user]
    (println "got a user " room " " user)
    (swap! users-per-room insert-user room user)
    (s/connect
      (bus/subscribe rooms room)
      socket)
    (bus/publish! rooms room
      (json/write-str {:type "room-update" :users (get @users-per-room room)})))

(defn new-connection-handler
  [request]
  (d/let-flow [socket (d/catch
                     (http/websocket-connection request)
                     (fn [_] nil))]
              (if-not socket
                non-websocket-request
                (d/let-flow [message-str (s/take! socket)]
		  (let [message (json/read-str message-str :key-fn keyword)]
		   (println message)
		    (if (= (get message :type) "join")
		      (let
		        [room (get message :room)
		         user (get message :user)]
	                (join-handler socket room user))
		      nil))))))

(def handler
  (params/wrap-params
   (compojure/routes
    ;; FIXME: We should decide if we want this route ("/") or another one
    (GET "/" [] new-connection-handler)
    (route/not-found "No such page."))))



(defn serve []
  (println (json/read-str "{\"test\": 10}" :key-fn keyword))
  (http/start-server handler
;; FIXME: We should decide the port to use
                     {:port 5000}))



