(ns racegex.websocket
  (:require [aleph.http :as http]
            [compojure.core :as compojure :refer [GET]]
            [compojure.route :as route]
            [manifold.bus :as bus]
            [manifold.deferred :as d]
            [manifold.stream :as s]
            [ring.middleware.params :as params]))

(def non-websocket-request
  {:status 400
   :headers {"content-type" "application/text"}
   :body "Expected a websocket request."})


(defn echo-handler
  [req]
  (->
   (d/let-flow [socket (http/websocket-connection req)]
               (s/connect socket socket))
   (d/catch
    (fn [_]
      non-websocket-request))))

(def chatrooms (bus/event-bus))

(defn chat-handler
  [req]
  (d/let-flow [conn (d/catch
                     (http/websocket-connection req)
                     (fn [_] nil))]
              (if-not conn
      ;; if it wasn't a valid websocket handshake, return an error
                non-websocket-request
      ;; otherwise, take the first two messages, which give us the chatroom and name
                ((d/let-flow [room (s/take! conn)
                              name (s/take! conn)]
      ;; take all messages from the chatroom, and feed them to the client
                             (s/connect
                              (bus/subscribe chatrooms room)
                              conn)
      ;; take all messages from the client, and publish it to the room
                             (s/consume
                              #(bus/publish! chatrooms room %)
                              (->> conn
                                   (s/map #(str %))
                                   (s/buffer 100)))
      ;; Compojure expects some sort of HTTP response, so just give it `nil`
                             nil)))))

(def handler
  (params/wrap-params
   (compojure/routes
    (GET "/echo" [] echo-handler)
    ;; FIXME: We should decide if we want this route ("/") or another one
    (GET "/" [] chat-handler)
    (route/not-found "No such page."))))



(defn serve []
  (http/start-server handler
;; FIXME: We should decide the port to use
                     {:port 5000}))



