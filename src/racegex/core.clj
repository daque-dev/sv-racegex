(require '[racegex.websocket])

(ns racegex.core
  (:gen-class))

(defn -main
  "Starts the server"
  [& args]
  (racegex.websocket/serve))
